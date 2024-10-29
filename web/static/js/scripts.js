// Fetch suggestions based on the user's input
function fetchSuggestions() {
    let query = document.getElementById("search-box").value;

    if (query.length === 0) {
        // Clear dropdown if the input is empty
        clearDropdown();
        return;
    }

    // Perform AJAX request to the Go backend
    fetch(`/suggestions?artist=${encodeURIComponent(query)}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => response.json())
    .then(data => populateDropdown(data))
    .catch(error => console.error('Error fetching suggestions:', error));
}

// Populate the dropdown with suggestions
function populateDropdown(suggestions) {
    let dropdown = document.getElementById("suggestions-dropdown");
    clearDropdown(); // Clear previous suggestions

    // make visible when dropdown has content
    if (suggestions.length > 0) {
        dropdown.style.display = 'block';
        dropdown.size = Math.min(suggestions.length, 5);
    }

    // Add each suggestion to the dropdown
    suggestions.forEach(suggestion => {
        let option = document.createElement("option");
        option.value = suggestion.Id;
        option.text = `${suggestion.QueryResult} - ${suggestion.SearchParam}`;
        dropdown.appendChild(option);
    });
}

// Clear the dropdown
function clearDropdown() {
    let dropdown = document.getElementById("suggestions-dropdown");
    dropdown.innerHTML = ""; // Clear all child options
    dropdown.size = 0; //Reset dropdown size
    dropdown.style.display = 'none'; // Hide dropdown menu when cleared
}

function redirectToArtist() {
    let dropdown = document.getElementById("suggestions-dropdown");
    let selectedArtistId = dropdown.value

    if (selectedArtistId) {
        window.location.href = `/artist/?id=${selectedArtistId}`;
    }
}

document.getElementById("suggestions-dropdown").addEventListener('click', redirectToArtist);