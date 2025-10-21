import { GrLogout } from "react-icons/gr";
import { CgProfile } from "react-icons/cg";
import styles from "./UserDropdown.module.css";

function UserDropdown() {
	return (
		<div>
			<ul className={styles.userDropList}>
				<div className={styles.dropOption}>
					<CgProfile />
					<li>Profile</li>
				</div>

				<div className={styles.dropOption}>
					<GrLogout />
					<li>Logout</li>
				</div>
			</ul>
		</div>
	);
}

export default UserDropdown;
