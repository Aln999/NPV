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

    // Call the function to visualize the network flow
    visualizeNetworkFlow(analysis);
}

function visualizeNetworkFlow(analysis) {
    const visualizationDiv = document.getElementById('network-flow-visualization');
    visualizationDiv.innerHTML = ''; // Clear previous visualization

    // Example: Create a simple representation of network flow
    const svg = d3.select(visualizationDiv)
        .append('svg')
        .attr('width', 600)
        .attr('height', 400);

    // Here you can customize the visualization based on the analysis data
    // For example, drawing circles for pods and lines for connections
    // This is a placeholder for your actual visualization logic
    svg.append('circle')
        .attr('cx', 100)
        .attr('cy', 200)
        .attr('r', 30)
        .style('fill', 'blue');

    svg.append('text')
        .attr('x', 100)
        .attr('y', 200)
        .attr('dy', '.35em')
        .text('Pod A');

    svg.append('circle')
        .attr('cx', 300)
        .attr('cy', 200)
        .attr('r', 30)
        .style('fill', 'green');

    svg.append('text')
        .attr('x', 300)
        .attr('y', 200)
        .attr('dy', '.35em')
        .text('Pod B');

    svg.append('line')
        .attr('x1', 130)
        .attr('y1', 200)
        .attr('x2', 270)
        .attr('y2', 200)
        .attr('stroke', 'black');
}
