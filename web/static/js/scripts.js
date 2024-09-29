document.getElementById('search-box').addEventListener('input', function() {
    const query = this.value;
    if (query.length === 0) {
        document.getElementById('suggestions').style.display = 'none';
        return;
    }

    fetch('/results?q=' + encodeURIComponent(query))
    .then(response => response.json())
    .then(data => {
        const suggestions = document.getElementById('suggestions');
        suggestions.innerHTML = '';

        if (data.length === 0) {
            suggestions.style.display = 'none';
            return;
        }

        data.forEach(artist => {
            const li = document.createElement('li');
            li.textContent = artist.name;
            li.onclick = function() {
                document.getElementById('search-box').value = artist.name;
                suggestions.style.display = 'none';
                // Add your code to display artist details here
            };
            suggestions.appendChild(li);
        });

        suggestions.style.display = 'block';
    });
});