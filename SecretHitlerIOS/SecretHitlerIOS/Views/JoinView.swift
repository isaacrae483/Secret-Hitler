//
// Created by Isaac Rae on 2/3/23.
//

import Foundation
import SwiftUI

struct JoinView: View {

    @StateObject var viewModel: ViewModel = ViewModel()

    var body: some View {
        VStack {
            Text("Secret Hitler")
                    .font(.title)

            Button("Join Room") {
                viewModel.joinRoom()
            }
                    .buttonStyle(.rounded)
            Button("Create Room") {
                viewModel.createRoom()
            }
                    .buttonStyle(.rounded)
            Button("Browse Rooms") {
                viewModel.browseRoom()
            }
                    .buttonStyle(.rounded)
        }
                .sheet(isPresented: $viewModel.showSheet) {
                    VStack {
                        TextField("Room Code", text: $viewModel.code)
                        Button("Join") {
                            viewModel.showSheet = false
                        }
                                .buttonStyle(.rounded)
                    }
                            .padding(.horizontal, 14)
                }
    }
}
