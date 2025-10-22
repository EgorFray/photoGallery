import { useAuth } from "../context/FakeAuthContext";
import styles from "./UserContent.module.css";

function UserContent() {
	const { user } = useAuth();

	return (
		<div className={styles.userInfo}>
			<img
				src={`${import.meta.env.VITE_BACKEND_URL}${user.avatar}`}
				alt={user.name}
				className={styles.userImg}
			/>
			<div className={styles.userInfoRight}>
				<p>{user.name}</p>
				<p>{user.email}</p>
			</div>
		</div>
	);
}

export default UserContent;
