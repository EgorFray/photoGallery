import Footer from "../components/Footer";
import Layout from "../components/Layout";
import Main from "../components/Main";
import NavBar from "../components/NavBar";
import PostContent from "../components/PostContent";

function PostDetail() {
	return (
		<Layout>
			<NavBar />
			<Main>
				<PostContent />
			</Main>
			<Footer />
		</Layout>
	);
}

export default PostDetail;
