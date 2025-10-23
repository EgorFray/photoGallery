import { Navigate, useNavigate } from "react-router-dom";
import { useAuth } from "../context/FakeAuthContext";
import Button from "./Button";
import styles from "./UserContent.module.css";

function UserContent() {
	const { user } = useAuth();

	const navigate = useNavigate();

	return (
		<>
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
			<Button
				className={styles.updateBtn}
				onClick={() => navigate("/profile/update")}
			>
				Update
			</Button>
		</>
	);
}

export default UserContent;
