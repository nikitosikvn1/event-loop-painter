const form = document.querySelector('form');
const textarea = document.querySelector('textarea');

form.addEventListener('submit', (e) => {
    e.preventDefault();
    const params = textarea.value.trim().split('\n');
    const queryParams = params.map(param => encodeURIComponent(param)).join(',');
    const url = `http://127.0.0.1:17000/?cmd=${queryParams}`;

    console.log(url);

    fetch(url)
    .then(response => {
        if (!response.ok) throw new Error('Network response was not ok');
        return response.text();
    })
    .then(data => {
        console.log(data);
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });
});
