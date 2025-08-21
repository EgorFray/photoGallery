import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/FakeAuthContext";
import styles from "./User.module.css";

function User() {
	const { user, logout } = useAuth();
	const navigate = useNavigate();

	function handleClick() {
		logout();
		navigate("/login");
	}

	return (
		<div className={styles.user}>
			<img src={user.avatar} alt={user.name} />
			<button onClick={handleClick}>Logout</button>
		</div>
	);
}

export default User;
