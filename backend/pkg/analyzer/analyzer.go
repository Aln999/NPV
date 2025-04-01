package analyzer

import (
	"context"

	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PolicyAnalysis struct {
	WideOpenPolicies    []string
	UnnecessaryPolicies []string
	Recommendations     []string
}

func AnalyzePolicies(clientset *kubernetes.Clientset, namespace string) (PolicyAnalysis, error) {
	policies, err := clientset.NetworkingV1().NetworkPolicies(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return PolicyAnalysis{}, err
	}

	var analysis PolicyAnalysis

	for _, policy := range policies.Items {
		if isPolicyWideOpen(policy) {
			analysis.WideOpenPolicies = append(analysis.WideOpenPolicies, policy.Name)
		}

		if isPolicyUnnecessary(policy) {
			analysis.UnnecessaryPolicies = append(analysis.UnnecessaryPolicies, policy.Name)
		}
	}

	analysis.Recommendations = []string{
		"Restrict wide-open policies using specific pod selectors",
		"Remove policies allowing all ingress/egress traffic",
	}

	return analysis, nil
}

func isPolicyWideOpen(policy netv1.NetworkPolicy) bool {
	// Check if policy allows all ingress
	for _, ingress := range policy.Spec.Ingress {
		if len(ingress.From) == 0 {
			return true
		}
	}
	return false
}

func isPolicyUnnecessary(policy netv1.NetworkPolicy) bool {
	// Check if policy has no matching pods
	return len(policy.Spec.PodSelector.MatchLabels) == 0
}
