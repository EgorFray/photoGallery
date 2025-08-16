import { useState } from "react";
import Footer from "../components/Footer";
import Header from "../components/Header";
import Layout from "../components/Layout";
import List from "../components/List";
import Main from "../components/Main";
import NavBar from "../components/NavBar";

function MainPage() {
	return (
		<div>
			<Layout>
				<NavBar />
				<Header />
				<Main>
					<List />
				</Main>
				<Footer />
			</Layout>
		</div>
	);
}

export default MainPage;
