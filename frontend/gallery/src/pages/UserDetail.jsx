import Footer from "../components/Footer";
import Layout from "../components/Layout";
import Main from "../components/Main";
import NavBar from "../components/NavBar";
import UserContent from "../components/UserContent";

function UserDetail() {
	return (
		<Layout>
			<NavBar />
			<Main>
				<UserContent />
			</Main>
			<Footer />
		</Layout>
	);
}

export default UserDetail;
