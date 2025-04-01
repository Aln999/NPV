package recommender

import "fmt"

func SuggestNetworkPolicy(source, destination string) string {
	return fmt.Sprintf(`apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-%s-to-%s
spec:
  podSelector:
    matchLabels:
      app: %s
  policyTypes:
  - Ingress
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: %s`, destination, source, destination, source)
}
