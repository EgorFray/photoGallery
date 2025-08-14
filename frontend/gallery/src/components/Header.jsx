import styles from "./Header.module.css";
import { NavLink } from "react-router-dom";

function Header() {
	return (
		<div className={styles.header}>
			<NavLink to={"/"} className={styles.postsLink}>
				Posts
			</NavLink>

			<NavLink to={"/create"} className={styles.postsLink}>
				Add post
			</NavLink>
		</div>
	);
}

export default Header;
