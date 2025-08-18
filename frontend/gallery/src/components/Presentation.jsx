import Button from "./Button";
import styles from "./Presentation.module.css";

function Presentation() {
	return (
		<div className={styles.container}>
			<div className={styles.textBox}>
				<p className={styles.text}>Create your own beautifull gallery</p>
			</div>
			<div className={styles.imgBox}>
				<div className={styles.imgWrapper}>
					<img src="/green.jpg" />
				</div>
				<div className={styles.imgWrapper}>
					<img src="/mist.jpg" />
				</div>
				<div className={styles.imgWrapper}>
					<img src="/tent.jpg" />
				</div>
			</div>
			<Button>Login</Button>
		</div>
	);
}

export default Presentation;
