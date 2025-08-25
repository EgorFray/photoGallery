import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/FakeAuthContext";
import Button from "./Button";
import styles from "./LoginForm.module.css";

function LoginForm() {
	const [email, setEmail] = useState("jackexample@gmail.com");
	const [password, setPassword] = useState("qwerty");

	const { login, isAuthenticated } = useAuth();
	const navigate = useNavigate();

	function handleSubmit(e) {
		e.preventDefault();

		if (email && password) login(email, password);
	}

	useEffect(
		function () {
			if (isAuthenticated) navigate("/app", { replace: true });
		},
		[isAuthenticated, navigate]
	);

	if (isAuthenticated) return null;

	return (
		<div className={styles.formContainer}>
			<form className={styles.form} onSubmit={handleSubmit}>
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

				<Button className={styles.btnForm}>Login</Button>
			</form>
		</div>
	);
}

export default LoginForm;
