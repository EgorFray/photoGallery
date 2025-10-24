import { useActionData, useLocation } from "react-router-dom";
import { useAuth } from "../context/FakeAuthContext";
import styles from "./NavBar.module.css";
import Search from "../components/Search";
import Logo from "./Logo";
import Empty from "./Empty";
import SearchPlaceholder from "./SearchPlaceholder";
import User from "./User";

function NavBar() {
	const location = useLocation();
	const isMainPage = location.pathname === "/app";
	const { isAuthenticated } = useAuth();

	return (
		<nav className={styles.navbar}>
			<Logo />
			{/* We neeed this to correctly show borders in the NavBar while <Search /> is hidden */}
			{isMainPage ? <Search /> : <SearchPlaceholder />}
			{isAuthenticated ? <User /> : <Empty />}
		</nav>
	);
}

export default NavBar;
