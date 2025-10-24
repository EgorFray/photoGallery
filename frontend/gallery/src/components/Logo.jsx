import { Link } from "react-router-dom";
import styles from "./Logo.module.css";

function Logo() {
	return (
		<div className={styles.logoBox}>
			<Link to="/app">
				<img src="/Logo.png" alt="logo of memorize" className={styles.logo} />
			</Link>
		</div>
	);
}

export default Logo;
