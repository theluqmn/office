const currentPath = window.location.pathname;

const navHomeKey = document.getElementById('nav-home-key');
const navProjectsKey = document.getElementById('nav-projects-key');
const navJournalKey = document.getElementById('nav-journal-key');
const navSitemapKey = document.getElementById('nav-sitemap-key');

const navHomeText = document.getElementById('nav-home-text');
const navProjectsText = document.getElementById('nav-projects-text');
const navJournalText = document.getElementById('nav-journal-text');
const navSitemapText = document.getElementById('nav-sitemap-text');

switch (currentPath) {
    case '/': navHomeText.classList.add('underline'); break
    case '/projects': navProjectsText.classList.add('underline'); break
}

document.addEventListener('keydown', (event) => {
    if (event.repeat) return;
    const targetTag = event.target.tagName.toLowerCase()
    if (targetTag === 'input' || targetTag === 'textarea') return;

    switch (event.key) {
        case '1': navHomeKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
        case '2': navProjectsKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
        case '3': navJournalKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
        case '4': navSitemapKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
    }
})

document.addEventListener('keyup', (event) => {
    if (event.repeat) return;
    switch (event.key) {
        case '1': navHomeKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/'; break;
        case '2': navProjectsKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/projects'; break;
        case '3': navJournalKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/'; break;
        case '4': navSitemapKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/'; break;
    }
});