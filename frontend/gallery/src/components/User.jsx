import { useEffect } from "react";
import { useState } from "react";
import { useUser } from "../context/UserContext";
import styles from "./User.module.css";
import UserDropdown from "./UserDropdown";

function User() {
	const [isOpenDrop, setIsOpenDrop] = useState(false);
	const { curUser, getCurrentUser } = useUser();

	useEffect(function () {
		getCurrentUser();
	}, []);

	return (
		<>
			<div className={styles.user}>
				<img
					src={`${import.meta.env.VITE_BACKEND_URL}${curUser.avatar}`}
					alt={curUser.name}
					onClick={() => setIsOpenDrop(!isOpenDrop)}
				/>
			</div>
			{isOpenDrop && <UserDropdown />}
		</>
	);
}

export default User;
