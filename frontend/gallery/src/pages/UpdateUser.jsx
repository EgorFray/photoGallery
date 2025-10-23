import Footer from "../components/Footer";
import Layout from "../components/Layout";
import Main from "../components/Main";
import NavBar from "../components/NavBar";
import UpdateUserForm from "../components/UpdateUserForm";

function UpdateUser() {
	return (
		<Layout>
			<NavBar />
			<Main>
				<UpdateUserForm />
			</Main>
			<Footer />
		</Layout>
	);
}

export default UpdateUser;
