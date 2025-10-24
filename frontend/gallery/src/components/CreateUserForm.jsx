import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useUser } from "../context/UserContext";
import Button from "./Button";
import styles from "./CreateUserForm.module.css";
import Spinner from "./Spinner";

function CreateUserForm() {
	const { createUser, isLoading } = useUser();
	const [email, setEmail] = useState("");
	const [name, setName] = useState("");
	const [password, setPassword] = useState("");

	const navigate = useNavigate();

	async function handleSubmit(e) {
		e.preventDefault();

		const formData = new FormData(e.target);
		await createUser(formData);
		navigate("/login");
	}

	if (isLoading) return <Spinner />;

	return (
		<div className={styles.formContainer}>
			<form className={styles.createForm} onSubmit={handleSubmit}>
				<h2 className={styles.createUserHeading} />

				<div className={styles.row}>
					<label htmlFor="name">Name</label>
					<input
						type="text"
						id="name"
						name="name"
						onChange={(e) => setName(e.target.value)}
						value={name}
					/>
				</div>

				<div className={styles.row}>
					<label htmlFor="email">Email</label>
					<input
						type="email"
						id="email"
						name="email"
						onChange={(e) => setEmail(e.target.value)}
						value={email}
					/>
				</div>

				<div className={styles.row}>
					<label htmlFor="password">Password</label>
					<input
						type="password"
						id="password"
						name="password"
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

				<Button className={styles.btnForm} type="submit">
					Create user
				</Button>
			</form>
		</div>
	);
}

export default CreateUserForm;
