
let slideIndex = 0;
let slides = $(".mySlides");
let slideInterval;

function startSlideshow() {
    slideInterval = setInterval(() => {
        showSlides(++slideIndex);
    }, 2000);
}

function pauseSlideshow() {
    clearInterval(slideInterval);
}

function changeSlide(n) {
    pauseSlideshow();
    showSlides(slideIndex += n);
}

function showSlides(n) {
    let i;

    if (n >= slides.length) {
        slideIndex = 0;
    }

    if (n < 0) {
        slideIndex = slides.length - 1;
    }

    slides.hide(); // Masque toutes les diapositives
    startSlideshow(); // Redémarre le diaporama

    // Affiche toutes les diapositives simultanément
    slides.show();
}

$("#main").on("mouseover", pauseSlideshow);
$("#main").on("mouseout", startSlideshow);

window.onload = startSlideshow;
