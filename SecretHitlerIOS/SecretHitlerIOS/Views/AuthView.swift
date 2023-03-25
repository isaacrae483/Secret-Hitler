//
// Created by Isaac Rae on 3/22/23.
//

import Foundation
import SwiftUI

struct AuthView: View {
    @StateObject private var viewModel: ViewModel

    init(_ service: Service, _ authenticator: Authenticator) {
        _viewModel = StateObject(wrappedValue: ViewModel(service, authenticator))
    }

    var body: some View {
        VStack {
            TextField("Username", text: $viewModel.authentication.username)
                    .textFieldStyle(.rounded)
            TextField("Password", text: $viewModel.authentication.password)
                    .textFieldStyle(.rounded)

            Button("Login" ) {
                viewModel.login()
            }
                    .buttonStyle(.rounded)
            Button("Signup") {
                viewModel.signup()
            }
                    .buttonStyle(.rounded)

        }
    }
}