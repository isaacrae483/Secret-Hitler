//
// Created by Isaac Rae on 3/22/23.
//

import Foundation
import SwiftUI
import Combine

extension AuthView {
    @MainActor class ViewModel: ObservableObject {
        @Published var username: String = "isaac"
        @Published var password: String = "hello_world"
        let authenticator: Authenticator
        let service: Service

        init(_ service: Service, _ authenticator: Authenticator) {
            self.service = service
            self.authenticator = authenticator
        }

        var cancelBag = Set<AnyCancellable>()

        func login() {
            service.auth.login(auth: Authentication(username: username, password: password))
                    .mapError({ (error) -> Error in // 5
                        print(error)
                        return error
                    })
                    .sink(receiveCompletion: { _ in }, // 6
                            receiveValue: {
                                print($0)
                                self.authenticator.login()
                            })
                    .store(in: &cancelBag)


        }

    }
}