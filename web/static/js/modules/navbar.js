export class Navbar {
	static init() {
		console.info("Initialize Navbar");

		const navbar = document.querySelector(".navbar");

		if (!navbar) {
			return;
		}

		const navItemsToggler = navbar.querySelector(".nav-items-toggler");
		const navItems = navbar.querySelector(".nav-items");

		if (!navItemsToggler || !navItems) {
			return;
		}

		navItemsToggler.addEventListener("click", () => {
			console.debug("Toggle Navigation Items");
			navItemsToggler.classList.toggle("change");
			navItems.classList.toggle("show");
		});
	}
}
