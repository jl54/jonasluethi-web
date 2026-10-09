export class Navbar {
    init() {
        const navbar = document.querySelector("header");

        if (!navbar) {
            return;
        }

        const navbarWrapper = navbar.parentNode;

        if (!navbarWrapper) {
            return;
        }

        this._checkIntersection(navbarWrapper, navbar);
    }

    _checkIntersection(navbarWrapper, navbar) {
        window.addEventListener("scroll", () => {
            if (navbarWrapper.getBoundingClientRect().top <= 0) {
                navbar.classList.add('is-sticky');
            } else {
                navbar.classList.remove('is-sticky');
            }
        }, {
            passive: true
        });
    }
}
