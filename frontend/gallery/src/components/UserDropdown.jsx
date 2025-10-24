import { GrLogout } from "react-icons/gr";
import { CgProfile } from "react-icons/cg";
import { motion } from "motion/react";
import { useAuth } from "../context/FakeAuthContext";
import { useNavigate } from "react-router-dom";
import { useUser } from "../context/UserContext";
import styles from "./UserDropdown.module.css";

function UserDropdown() {
	const { user, logout } = useAuth();
	const { getCurrentUser } = useUser();
	const navigate = useNavigate();

	function handleClickProfile() {
		getCurrentUser(user.Id);
		navigate("/profile");
	}

	function handleClickLogout() {
		logout();
		navigate("/login");
	}

	return (
		<motion.ul
			className={styles.userDropList}
			initial={{ opacity: 0, y: -10 }}
			animate={{ opacity: 1, y: 0 }}
			exit={{ opacity: 0, y: -10 }}
			transition={{ duration: 0.25 }}
		>
			<div className={styles.dropOption}>
				<CgProfile />
				<li onClick={handleClickProfile}>Profile</li>
			</div>

			<div className={styles.dropOption}>
				<GrLogout />
				<li onClick={handleClickLogout}>Logout</li>
			</div>
		</motion.ul>
	);
}

export default UserDropdown;
