const currentPath = window.location.pathname;

const navHomeKey = document.getElementById('nav-home-key');
const navJournalKey = document.getElementById('nav-journal-key');
const navToolsKey = document.getElementById('nav-tools-key');
const navSitemapKey = document.getElementById('nav-sitemap-key');

const navHomeText = document.getElementById('nav-home-text');
const navJournalText = document.getElementById('nav-journal-text');
const navToolsText = document.getElementById('nav-tools-text');
const navSitemapText = document.getElementById('nav-sitemap-text');

switch (currentPath) {
    case '/': navHomeText.classList.add('underline'); break
    
}

document.addEventListener('keydown', (event) => {
    if (event.repeat) return;
    const targetTag = event.target.tagName.toLowerCase()
    if (targetTag === 'input' || targetTag === 'textarea') return;

    switch (event.key) {
        case '1': navHomeKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
        case '2': navJournalKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
        case '3': navToolsKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
        case '4': navSitemapKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
    }
})

document.addEventListener('keyup', (event) => {
    if (event.repeat) return;
    switch (event.key) {
        case '1': navHomeKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/'; break;
        case '2': navJournalKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/'; break;
        case '3': navToolsKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/'; break;
        case '4': navSitemapKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/'; break;
    }
});