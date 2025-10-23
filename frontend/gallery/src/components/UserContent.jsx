import { Navigate, useNavigate } from "react-router-dom";
import { useAuth } from "../context/FakeAuthContext";
import { useUser } from "../context/UserContext";
import Button from "./Button";
import styles from "./UserContent.module.css";

function UserContent() {
	const { user } = useAuth();
	const { curUser } = useUser();

	const navigate = useNavigate();
	console.log(curUser);

	return (
		<>
			<div className={styles.userInfo}>
				<img
					src={`${import.meta.env.VITE_BACKEND_URL}${curUser.avatar}`}
					alt={user.name}
					className={styles.userImg}
				/>
				<div className={styles.userInfoRight}>
					<p>{curUser.name}</p>
					<p>{curUser.email}</p>
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
