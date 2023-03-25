//
// Created by Isaac Rae on 3/24/23.
//

import Foundation

class Authenticator : ObservableObject {
    @Published var authenticated: Bool = false

    func login() {
        authenticated = true
    }

    func logout() {
        authenticated = false
    }
}
