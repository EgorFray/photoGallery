import { useState } from "react";
import { useAuth } from "../context/FakeAuthContext";
import styles from "./User.module.css";
import UserDropdown from "./UserDropdown";

function User() {
	const [isOpenDrop, setIsOpenDrop] = useState(false);
	const { user } = useAuth();

	return (
		<>
			<div className={styles.user}>
				<img
					src={`${import.meta.env.VITE_BACKEND_URL}${user.avatar}`}
					alt={user.name}
					onClick={() => setIsOpenDrop(!isOpenDrop)}
				/>
			</div>
			{isOpenDrop && <UserDropdown />}
		</>
	);
}

export default User;
