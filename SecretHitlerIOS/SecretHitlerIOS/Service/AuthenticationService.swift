//
// Created by Isaac Rae on 3/21/23.
//

import Foundation
import Combine

class AuthenticationService {
    let client: APIClient

    init(_ client: APIClient) {
        self.client = client
    }

    func login(auth: Authentication) -> AnyPublisher<[String: String], NetworkServiceError> {
        client.request(.login, "POST", body: auth)
    }

    func signup(auth: Authentication) -> AnyPublisher<[String: String], NetworkServiceError> {
        client.request(.signup, "POST", body: auth)
    }
}