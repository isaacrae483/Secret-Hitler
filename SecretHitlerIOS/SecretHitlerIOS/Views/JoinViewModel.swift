//
// Created by Isaac Rae on 2/4/23.
//

import Foundation
import SwiftUI
import Combine

extension JoinView {
    @MainActor class ViewModel: ObservableObject {
        @Published var showSheet: Bool = false

        @Published var code: String = ""

        let service: Service

        init(_ service: Service) {
            self.service = service
        }

        func joinRoom() {
            print("joining")
            showSheet.toggle()
            print(showSheet)
        }

        func createRoom() {
            print("creating")
        }

        var cancellationToken: AnyCancellable? // 2

        func browseRoom() {
            cancellationToken = service.rooms.available()
                    .mapError({ (error) -> Error in // 5
                        print(error)
                        return error
                    })
                    .sink(receiveCompletion: { _ in }, // 6
                            receiveValue: {
                                print($0)
                            })
            print("browsing")
        }

    }
}