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
            Button(action: viewModel.login) {
                Text("login")
            }
        }
    }
}