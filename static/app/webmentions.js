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
                li.setAttribute('class', 'webmentions__item');
                const a = document.createElement('a');
                a.setAttribute('href', mention.source);
                if (mention.author_name && mention.type === 'comment') {
                    a.innerText = mention.author_name
                } else {
                    a.innerText = mention.source
                }
                const icon = document.createElement('i');
                if (mention.type === 'comment') {
                    icon.setAttribute('class', 'far fa-comment');
                } else {
                    icon.setAttribute('class', 'far fa-link');
                }
                li.appendChild(icon);
                li.appendChild(a);
                if (mention.created_at) {
                    const createdAt = document.createElement('span');
                    createdAt.setAttribute('class', 'webmentions__item__date');
                    createdAt.innerText = '@ ' + mention.created_at;
                    li.appendChild(createdAt);
                }
                if (mention.content) {
                    const bq = document.createElement('blockquote');
                    bq.innerText = mention.content;
                    bq.setAttribute('class', 'webmentions__item__comment');
                    li.appendChild(bq);
                }
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
