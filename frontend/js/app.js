async function checkConnectivity() {
    const source = document.getElementById('source').value;
    const destination = document.getElementById('destination').value;
    
    const response = await fetch(`/api/check?source=${encodeURIComponent(source)}&destination=${encodeURIComponent(destination)}`);
    const result = await response.json();
    
    const resultDiv = document.getElementById('connectivity-result');
    resultDiv.innerHTML = `
        <h3>Connectivity Status: ${result.connectivity}</h3>
        ${result.recommendations ? `
        <h4>Recommendations:</h4>
        <pre>${result.recommendations}</pre>
        ` : ''}
    `;
}

async function analyzePolicies() {
    const namespace = document.getElementById('namespace').value;
    
    const response = await fetch(`/api/analyze?namespace=${encodeURIComponent(namespace)}`);
    const analysis = await response.json();
    
    const resultDiv = document.getElementById('analysis-result');
    resultDiv.innerHTML = `
        <h3>Analysis Results:</h3>
        <p>Wide-Open Policies: ${analysis.wideOpenPolicies.join(', ') || 'None'}</p>
        <p>Unnecessary Policies: ${analysis.unnecessaryPolicies.join(', ') || 'None'}</p>
        <h4>Recommendations:</h4>
        <ul>${analysis.recommendations.map(r => `<li>${r}</li>`).join('')}</ul>
    `;
}