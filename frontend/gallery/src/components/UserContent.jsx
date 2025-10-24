import { Navigate, useNavigate } from "react-router-dom";
import { useAuth } from "../context/FakeAuthContext";
import { useUser } from "../context/UserContext";
import Button from "./Button";
import styles from "./UserContent.module.css";

function UserContent() {
	const { user } = useAuth();
	const { curUser } = useUser();

	const navigate = useNavigate();

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

			<div className={styles.btnWrapper}>
				<Button onClick={() => navigate("/profile/update")}>Update</Button>
			</div>
		</>
	);
}

export default UserContent;
