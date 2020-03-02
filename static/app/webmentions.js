(function() {
    const baseURL = "https://zerokspot.com";
    const targetURL = `${baseURL}${window.location.pathname}`;
    const apiEndpoint = `https://zerokspot.com/webmentions/get?target=${encodeURIComponent(targetURL)}`;
    return fetch(apiEndpoint).then(resp => {
        return resp.json().then(mentions => {
            if (mentions.length === 0) {
                return;
            }
            const container = document.querySelector(".webmentions");
            const fragment = document.createDocumentFragment();
            const title = document.createElement('h2');
            const listing = document.createElement('ul');
            title.innerText = `Webmentions (${mentions.length})`;
            title.className = 'webmentions__title';
            mentions.forEach(mention => {
                const li = document.createElement('li');
                const a = document.createElement('a');
                a.setAttribute('href', mention.source);
                a.innerText = mention.source
                li.appendChild(a);
                listing.appendChild(li);
            });
            fragment.appendChild(title);
            fragment.appendChild(listing);
            container.appendChild(fragment);
        });
    }, err => {
        console.log(err);
    });
})()
