import styles from "./NavBar.module.css";
import Search from "../components/Search";
import Logo from "./Logo";
import Empty from "./Empty";
import { useLocation } from "react-router-dom";
import SearchPlaceholder from "./SearchPlaceholder";

function NavBar() {
	const location = useLocation();
	const isMainPage = location.pathname === "/app";
	return (
		<nav className={styles.navbar}>
			<Logo />
			{/* We neeed this to correctly show borders in the NavBar while <Search /> is hidden */}
			{isMainPage ? <Search /> : <SearchPlaceholder />}
			<Empty />
		</nav>
	);
}

export default NavBar;
