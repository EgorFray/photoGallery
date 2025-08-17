import Footer from "../components/Footer";
import Layout from "../components/Layout";
import Main from "../components/Main";
import NavBar from "../components/NavBar";
import NotFound from "../components/NotFound";

function PageNotFound() {
	return (
		<Layout>
			<NavBar />
			<Main>
				<NotFound />
			</Main>
			<Footer />
		</Layout>
	);
}

export default PageNotFound;
