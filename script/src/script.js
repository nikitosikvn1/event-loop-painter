const form = document.querySelector('form');
const textarea = document.querySelector('textarea');
const gfButton = document.querySelector('.gf-script');
const dmButton = document.querySelector('.dm-script');

async function sendHTTPRequest(url) {
    try {
        const response = await fetch(url);
        if (response.ok) return response.text();
        throw new Error(`Request failed with status ${response.status}`);
    } catch {
        console.error(error);
    }
}

form.addEventListener('submit', (e) => {
    e.preventDefault();
    const params = textarea.value.trim().split('\n');
    const queryParams = params.map(param => encodeURIComponent(param)).join(',');
    const url = `http://127.0.0.1:17000/?cmd=${queryParams}`;

    sendHTTPRequest(url)
      .then(response => console.log(response))
      .catch(error => console.error(error));
});

gfButton.addEventListener('click', () => {
    const url = 'http://127.0.0.1:17000/?cmd=green,bgrect%200.05%200.05%200.95%200.95,update';

    sendHTTPRequest(url)
      .then(response => console.log(response))
      .catch(error => console.error(error));
});

dmButton.addEventListener('click', () => {
    const urlToDraw = 'http://127.0.0.1:17000/?cmd=white,figure%200.1%200.1,update';
    const urlToMove = 'http://127.0.0.1:17000/?cmd=move%200.1%200.1,update';

    sendHTTPRequest(urlToDraw);

    for (let i = 0; i < 9; i++) {
        setTimeout(() => {
            console.log(`Request: ${i + 1}`);
            sendHTTPRequest(urlToMove);
        }, (i + 1) * 1000);
    }
});