/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ratelimitv1 "ratelimiteroperator/api/v1"
	rateservice "ratelimiteroperator/grpc"
)

// RateLimitPolicyReconciler reconciles a RateLimitPolicy object
type RateLimitPolicyReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	grpcConn   *grpc.ClientConn
	grpcClient rateservice.RateServiceClient
}

// +kubebuilder:rbac:groups=ratelimit.ratelimit.io,resources=ratelimitpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ratelimit.ratelimit.io,resources=ratelimitpolicies/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ratelimit.ratelimit.io,resources=ratelimitpolicies/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RateLimitPolicy object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.4/pkg/reconcile
func (r *RateLimitPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the RateLimitPolicy instance
	var policyList ratelimitv1.RateLimitPolicyList
	if err := r.List(ctx, &policyList); err != nil {
		logger.Error(err, "unable to fetch RateLimitPolicy")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Apply rate limiting rules (dummy example)
	for _, policy := range policyList.Items {
		req := &rateservice.ServiceRateMessage{
			Service: policy.Spec.Service,
			Rpm:     int32(policy.Spec.MaxRequests),
		}

		resp, err := r.grpcClient.SetRate(ctx, req)
		if err != nil {
			logger.Error(err, "Failed to update rate limit for", policy.Name, err)
			continue
		}
		logger.Info("gRPC Response for ", policy.Name, resp)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RateLimitPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {

	grpcAddress := "ratelimiter-service.default.svc.cluster.local:3001"
	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	// Store gRPC client for reuse
	r.grpcConn = conn
	r.grpcClient = rateservice.NewRateServiceClient(conn)
	return ctrl.NewControllerManagedBy(mgr).
		For(&ratelimitv1.RateLimitPolicy{}).
		Named("ratelimitpolicy").
		Complete(r)
}
