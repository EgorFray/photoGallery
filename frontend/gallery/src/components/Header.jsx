import styles from "./Header.module.css";

function Header({ onOpen }) {
	return (
		<div className={styles.header}>
			<div className={styles.headerBox}>
				<div className={styles["headerBox--left"]}>
					<h1 className={styles["headerText--left"]}>So, what's new?</h1>
				</div>

				<div className={styles["headerBox--right"]}>
					<h2 className={styles["headerText--right"]}>Only your memories</h2>
					<button className={styles.postButton} onClick={onOpen}>
						Post
					</button>
				</div>
			</div>
		</div>
	);
}

export default Header;
