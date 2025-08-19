import Button from "./Button";
import styles from "./LoginForm.module.css";

function LoginForm() {
	return (
		<form className={styles.form}>
			<div className={styles.row}>
				<label htmlFor="email">Email</label>
				<input type="email" id="email" />
			</div>

			<div className={styles.row}>
				<label htmlFor="password">Password</label>
				<input type="password" id="password" />
			</div>

			<Button>Login</Button>
		</form>
	);
}

export default LoginForm;
