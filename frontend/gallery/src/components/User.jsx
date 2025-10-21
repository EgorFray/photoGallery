import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/FakeAuthContext";
import styles from "./User.module.css";
import UserDropdown from "./UserDropdown";

function User() {
	const { user, logout } = useAuth();
	const navigate = useNavigate();

	return (
		<>
			<div className={styles.user}>
				<img
					src={`${import.meta.env.VITE_BACKEND_URL}${user.avatar}`}
					alt={user.name}
				/>
			</div>
			<UserDropdown />
		</>
	);
}

export default User;
