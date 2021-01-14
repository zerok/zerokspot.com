var zs = {icons: {}};
function renderIcons() {
  function getIconConfig(classes) {
    var result = {
      icon: null,
      group: 'regular'
    };
    classes.forEach(cls => {
      if (cls.startsWith('icongroup-')) {
        result.group = cls.substr(10);
        return;
      }
      if (!cls.startsWith('icon-')) {
        return;
      }
      result.icon = cls.substr(5);
    });
    return result;
  }
  document.querySelectorAll('.icon').forEach((elem) => {
    var ns = 'http://www.w3.org/2000/svg';
    var cfg = getIconConfig(elem.classList);
    if (!cfg.icon) {
      return;
    }
    var svg = zs.icons.createIconElement(cfg);
    elem.parentNode.replaceChild(svg, elem);
  });
}

zs.icons.createIconElement = function(cfg) {
  var ns = 'http://www.w3.org/2000/svg';
  var svg = document.createElementNS(ns, 'svg');
  var icon = window.icons[cfg.group][cfg.icon];
  var path = document.createElementNS(ns, 'path');
  path.setAttributeNS(null, 'd', icon.data);
  path.setAttributeNS(null, 'fill', 'currentColor');
  svg.appendChild(path);
  svg.classList.add('icon');
  svg.setAttribute('xmlns', 'http://www.w3.org/2000/svg');
  svg.setAttributeNS(null, 'viewBox', icon.viewBox.join(' '));
  svg.setAttribute('aria-hidden', 'true');
  svg.setAttribute('role', 'img');
  return svg;
}

window.zs = zs;

(function() {
    var button = document.createElement('a');
    function toggleDarkMode(goDark) {
      while(button.hasChildNodes()) {
        button.removeChild(button.firstChild);
      }
      if (goDark) {
        button.appendChild(window.zs.icons.createIconElement({group: 'regular', icon: 'sun'}));
        document.getElementsByTagName('html')[0].classList.add('darkmode');
      } else {
        button.appendChild(window.zs.icons.createIconElement({group: 'regular', icon: 'moon'}));
        document.getElementsByTagName('html')[0].classList.remove('darkmode');
      }
      window.localStorage.setItem('darkMode', goDark);
    }
    var goDark = window.localStorage.getItem('darkMode');
    if (goDark == null) {
        goDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
    } else {
        goDark = goDark === 'true';
    }
    var navList = document.querySelector('#main-nav > ul');
    var navItem = document.createElement('li');
    navItem.classList.add('darkmode-toggle');
    toggleDarkMode(goDark);
    button.setAttribute('href', '');
    navItem.appendChild(button);
    navList.appendChild(navItem);
    button.addEventListener('click', (evt) => {
        evt.preventDefault();
        goDark = !goDark;
        toggleDarkMode(goDark);
    }, false);
}());

function renderMastodonComments() {
  var container = document.querySelector('.mastodon-comments');
  if (!container) {
    return;
  }
  var baseURL = container.getAttribute('data-url');
  const list = document.createElement('div');
  list.className = 'mastodon-comments-container';
  fetch('https://zerokspot.com/retoots/api/v1/interactions?status=' + baseURL).then((resp) => {
    return resp.json();
  }).then(data => {
    const comments = data.descendants;
    const favoritedBy = data.favorited_by;
    if (comments.length) {
      comments.forEach(item => {
        var comment = document.createElement('div');
        comment.className = 'mastodon-comment';
        var avatar = document.createElement('div');
        avatar.className = 'mastodon-comment__avatar';
        if (item.account.avatar) {
          var avatarLink = document.createElement('a');
          avatarLink.setAttribute('href', item.account.url);
          var img = document.createElement('img');
          img.setAttribute('src', item.account.avatar);
          img.setAttribute('alt', item.account.username);
          avatarLink.appendChild(img);
          avatar.appendChild(avatarLink);
        }
        var content = document.createElement('div');
        content.className = 'mastodon-comment__content';
        content.innerHTML = item.content;
        var dateLink = document.createElement('a');
        dateLink.className = 'mastodon-comment__date';
        dateLink.setAttribute('href', item.url);
        dateLink.innerHTML = item.created_at;
        content.appendChild(dateLink);
        comment.appendChild(avatar);
        comment.appendChild(content);
        list.appendChild(comment)
      });
      container.appendChild(list);
    }
    if (favoritedBy.length) {
      const favoritedContainer = document.createElement('div');
      favoritedContainer.className = 'mastodon-favorited';
      const favoritedTitle = document.createElement('h3');
      favoritedTitle.innerText = 'Favorited by (' + favoritedBy.length + ')';
      favoritedContainer.appendChild(favoritedTitle);
      favoritedList = document.createElement('ul');
      favoritedContainer.appendChild(favoritedList);
      favoritedBy.forEach(item => {
        const li = document.createElement('li');
        li.className = 'mastodon-favorited__item';
        const link = document.createElement('a');
        link.setAttribute('href', item.url);
        const img = document.createElement('img');
        img.setAttribute('src', item.avatar);
        img.setAttribute('alt', item.username);
        link.appendChild(img);
        li.appendChild(link);
        favoritedList.appendChild(li);
      });
      container.appendChild(favoritedContainer);
    }
  });
}

renderIcons();
renderMastodonComments();
