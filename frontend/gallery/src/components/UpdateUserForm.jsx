import { motion } from "motion/react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/FakeAuthContext";
import { useUser } from "../context/UserContext";
import Button from "./Button";
import styles from "./UpdateUserForm.module.css";

function UpdateUserForm() {
	const { user } = useAuth();
	const { getCurrentUser, updateUser } = useUser();

	const [newName, setNewName] = useState(user.name);
	const [newPassword, setNewPassword] = useState("");

	const navigate = useNavigate();

	async function handleSubmit(e) {
		e.preventDefault();

		const formData = new FormData(e.target);
		await updateUser(formData);
		await getCurrentUser(user.Id);
		navigate("/profile");
	}

	return (
		<div className={styles.formContainer}>
			<p className={styles.updateHeading}>
				You don't have to update all field at once
			</p>
			<motion.span
				className={styles.updateHeadingMotion}
				initial={{ clipPath: "inset(0 100% 0 0)" }}
				animate={{ clipPath: "inset(0 0% 0 0)" }}
				transition={{ duration: 1.2, ease: "easeInOut", delay: 1.5 }}
			>
				Update only neccessary
			</motion.span>

			<form className={styles.updateForm} onSubmit={handleSubmit}>
				<div className={styles.row}>
					<label htmlFor="name">New name</label>
					<input
						type="text"
						id="name"
						name="name"
						onChange={(e) => setNewName(e.target.value)}
						value={newName}
					/>
				</div>

				<div className={styles.row}>
					<label htmlFor="password">New password</label>
					<input
						type="password"
						id="password"
						name="password"
						onChange={(e) => setNewPassword(e.target.value)}
						value={newPassword}
					/>
				</div>

				<div className={styles.row}>
					<label htmlFor="avatar" className={styles.addAvatar}>
						New avatar
					</label>
					<input type="file" name="avatar" id="avatar" />
				</div>
				<Button className={styles.btnForm} type="submit">
					Update user
				</Button>
			</form>
		</div>
	);
}

export default UpdateUserForm;
