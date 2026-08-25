const currentPage = document.body.id;
const activeNav = document.getElementById(`nav-${currentPage}-text`);

const navRepoKey = document.getElementById('nav-repo-key');
const navHomeKey = document.getElementById('nav-home-key');
const navProjectsKey = document.getElementById('nav-projects-key');
const navJournalKey = document.getElementById('nav-journal-key');

const navRepoText = document.getElementById('nav-repo-text');
const navHomeText = document.getElementById('nav-home-text');
const navProjectsText = document.getElementById('nav-projects-text');
const navJournalText = document.getElementById('nav-journal-text');

if (activeNav) {
    activeNav.classList.add('underline');
}

document.addEventListener('keydown', (event) => {
    if (event.repeat) return;
    const targetTag = event.target.tagName.toLowerCase()
    if (targetTag === 'input' || targetTag === 'textarea') return;

    switch (event.key) {
        case '0': navRepoKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
        case '1': navHomeKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
        case '2': navProjectsKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
        case '3': navJournalKey.classList.add('bg-[rgba(255,176,0,0.6)]'); break;
    }
})

document.addEventListener('keyup', (event) => {
    if (event.repeat) return;
    switch (event.key) {
        case '0': navRepoKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = 'https://github.com/theluqmn/office'; break;
        case '1': navHomeKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/'; break;
        case '2': navProjectsKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/projects'; break;
        case '3': navJournalKey.classList.remove('bg-[rgba(255,176,0,0.6)]'); window.location.href = '/journal'; break;
    }
});