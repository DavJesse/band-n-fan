// Fetch suggestions based on the user's input
function fetchSuggestions() {
    let query = document.getElementById("search-box").value;

    if (query.length === 0) {
        // Clear dropdown if the input is empty
        clearDropdown();
        return;
    }

    // Perform AJAX request to the Go backend
    fetch(`/suggestions?artist=${encodeURIComponent(query)}`)
        .then(response => response.json())
        .then(data => populateDropdown(data))
        .catch(error => console.error('Error fetching suggestions:', error));
}

// Populate the dropdown with suggestions
function populateDropdown(suggestions) {
    let dropdown = document.getElementById("suggestions-dropdown");
    clearDropdown(); // Clear previous suggestions

    // Hide dorpdown menu when empty
    if (suggestions.length === 0) {
        dropdown.style.display = 'none';
        return;
    }

    // make visible when dropdown has content
    dropdown.style.display = 'block';

    // Add each suggestion to the dropdown
    suggestions.forEach(suggestion => {
        let option = document.createElement("option");
        option.value = suggestion.id;
        option.text = suggestion.name;
        dropdown.appendChild(option);
    });
}

// Clear the dropdown
function clearDropdown() {
    let dropdown = document.getElementById("suggestions-dropdown");
    dropdown.innerHTML = ""; // Clear all child options
}

function redirectToArtist() {
    let dropdown = document.getElementById("suggestions-dropdown");
    let selectedArtistId = dropdown.value

    if (selectedArtistId) {
        window.location.href = `/artist/?id=${selectedArtistId}`;
    }
}

document.getElementById("suggestions-dropdown").addEventListener('click', redirectToArtist);