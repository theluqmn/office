canvas = document.getElementById("index-game");
ctx = canvas.getContext("2d");

document.fonts.ready.then(() => {
    ctx.font = "10px 'IBM Plex Mono'"
    ctx.fillStyle = "#ffb000";
    ctx.fillText("GELLO WORLD", 10, 50); 
})
