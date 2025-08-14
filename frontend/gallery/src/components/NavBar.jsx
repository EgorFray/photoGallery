import styles from "./NavBar.module.css";
import Search from "../components/Search";
import Logo from "./Logo";
import Empty from "./Empty";

function NavBar() {
	return (
		<nav className={styles.navbar}>
			<Logo />
			<Search />
			<Empty />
		</nav>
	);
}

export default NavBar;
