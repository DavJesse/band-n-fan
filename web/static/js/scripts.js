// Fetch suggestions based on the user's input
function fetchSuggestions() {
    let query = document.getElementById("search-box").value;
    console.log('Fetching suggestions for query:', query);

    if (query.length === 0) {
        // Clear dropdown if the input is empty
        clearDropdown();
        return;
    }

    // Perform AJAX request to the Go backend
    fetch(`/suggestions?artist=${encodeURIComponent(query)}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response not ok: '+ response.statusText)
            }
            return response.json();
        })
        .then(data => {
            console.log('Received suggestions: ', data);
            populateDropdown(data)

        })
        .catch(error => console.error('Error fetching suggestions:', error));
}

// Populate the dropdown with suggestions
function populateDropdown(suggestions) {
    let dropdown = document.getElementById("suggestions-dropdown");
    clearDropdown(); // Clear previous suggestions

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

function redirectToSearchResults() {
    const dropdown = document.getElementById("suggestions-dropdown");
    const selectedOption = dropdown.options[dropdown.selectIndex];

    if (selectedOption.value) {
        const artistId = selectedOption.value;
        window.location.href = `/artist/?id={artistId}`
    }
}