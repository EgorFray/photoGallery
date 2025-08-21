import styles from "./Logo.module.css";
import { Link } from "react-router-dom";

function Logo() {
	return (
		<div className={styles.logoBox}>
			<Link to="/app">
				<img src="/Logo.png" alt="" className={styles.logo} />
			</Link>
		</div>
	);
}

export default Logo;
