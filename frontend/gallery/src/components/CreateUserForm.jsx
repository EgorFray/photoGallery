import { useState } from "react";
import Button from "./Button";
import styles from "./CreateUserForm.module.css";

function CreateUserForm() {
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");

	function handleSubmit(e) {
		e.preventDefault();
	}

	return (
		<div className={styles.formContainer}>
			<form className={styles.createForm} onSubmit={handleSubmit}>
				<h2 className={styles.createUserHeading} />

				<div className={styles.row}>
					<label htmlFor="email">Email</label>
					<input
						type="email"
						id="email"
						onChange={(e) => setEmail(e.target.value)}
						value={email}
					/>
				</div>

				<div className={styles.row}>
					<label htmlFor="password">Password</label>
					<input
						type="password"
						id="password"
						onChange={(e) => setPassword(e.target.value)}
						value={password}
					/>
				</div>

				<div className={styles.row}>
					<label htmlFor="avatar" className={styles.addAvatar}>
						Add avatar
					</label>
					<input type="file" name="avatar" id="avatar" />
				</div>
			</form>

			<Button type="submit">Create user</Button>
		</div>
	);
}

export default CreateUserForm;
