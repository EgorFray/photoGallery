import { PostsProvider } from "./context/PostsContext";
import PostDetail from "./pages/PostDetail";
import { Routes, Route } from "react-router-dom";
import MainPage from "./pages/MainPage";
import CreatePost from "./pages/CreatePost";
import CreateUser from "./pages/CreateUser";
import PageNotFound from "./pages/PageNotFound";
import Homepage from "./pages/Homepage";
import Login from "./pages/Login";
import { AuthProvider } from "./context/FakeAuthContext";
import ProtectedRoute from "./pages/ProtectedRoute";
import { UserProvider } from "./context/UserContext";
import UserContent from "./components/UserContent";

function App() {
	return (
		<AuthProvider>
			<UserProvider>
				<PostsProvider>
					<Routes>
						<Route index element={<Homepage />} />
						<Route path="/login" element={<Login />} />
						<Route path="/createUser" element={<CreateUser />} />
						<Route
							path="/profile"
							element={
								<ProtectedRoute>
									<UserContent />
								</ProtectedRoute>
							}
						/>
						<Route
							path="/app"
							element={
								<ProtectedRoute>
									<MainPage />
								</ProtectedRoute>
							}
						/>
						<Route
							path="/create"
							element={
								<ProtectedRoute>
									<CreatePost />
								</ProtectedRoute>
							}
						/>
						<Route
							path="/app/posts/:id"
							element={
								<ProtectedRoute>
									<PostDetail />
								</ProtectedRoute>
							}
						/>
						<Route path="*" element={<PageNotFound />} />
					</Routes>
				</PostsProvider>
			</UserProvider>
		</AuthProvider>
	);
}

export default App;
