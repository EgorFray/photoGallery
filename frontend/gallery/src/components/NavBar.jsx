import styles from "./NavBar.module.css";

function NavBar() {
	return (
		<nav className={styles.navbar}>
			<div className={styles.logoBox}>
				<img src="../public/Logo.png" alt="" className={styles.logo} />
			</div>
			<div className={styles.textBox}>
				<p className={styles.memorize}>Memorize</p>
			</div>
		</nav>
	);
}

export default NavBar;
