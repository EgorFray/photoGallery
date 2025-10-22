import { useAuth } from "../context/FakeAuthContext";
import { useUser } from "../context/UserContext";
import styles from "./UserContent.module.css";

function UserContent() {
	return (
		<div className={styles.userInfo}>
			<img
				src={`${import.meta.env.VITE_BACKEND_URL}${user.avatar}`}
				alt={user.name}
			/>
			<div className={styles.userInfoRight}>
				<p>{user.name}</p>
				<p>{user.email}</p>
			</div>
		</div>
	);
}

export default UserContent;
